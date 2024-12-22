package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"

	"net"
	"strings"

	"github.com/Dushyantbha012/GO-IMDB/store"
)

// Start initializes the TCP server and handles incoming connections.
func Start(address string, s *store.Store) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn, s)
	}
}

func handleConnection(conn net.Conn, s *store.Store) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		command := strings.TrimSpace(line)
		response := handleCommand(conn, command, s)
		conn.Write([]byte(response + "\n"))
	}
}

func handleCommand(conn net.Conn, command string, s *store.Store) string {
	parts := strings.Fields(command)
	if len(parts) < 1 {
		return "ERR: Invalid command"
	}

	switch strings.ToUpper(parts[0]) {
	case "SET":
		if len(parts) < 3 {
			return "ERR: Usage SET <key> <value>"
		}
		key := parts[1]
		// Join remaining parts and handle quoted strings
		value := strings.Join(parts[2:], " ")
		if len(value) >= 2 && value[0] == '"' && value[len(value)-1] == '"' {
			// Remove quotes and unescape the string
			value = value[1 : len(value)-1]
			value = strings.ReplaceAll(value, `\"`, `"`)
		}
		s.Set(key, value)
		return "OK"

	case "GET":
		if len(parts) != 2 {
			return "ERR: Usage GET <key>"
		}
		if value, exists := s.Get(parts[1]); exists {
			return value
		}
		return "(nil)"

	case "LPUSH":
		if len(parts) < 3 {
			return "ERR: Usage LPUSH <key> <value> [<value>...]"
		}
		key := parts[1]
		values := parts[2:]
		l := s.GetOrCreateList(key)
		newLen := l.LPush(values...)
		return fmt.Sprintf("%d", newLen)

	case "RPUSH":
		if len(parts) < 3 {
			return "ERR: Usage RPUSH <key> <value> [<value>...]"
		}
		key := parts[1]
		values := parts[2:]
		l := s.GetOrCreateList(key)
		newLen := l.RPush(values...)
		return fmt.Sprintf("%d", newLen)

	case "LPOP":
		if len(parts) != 2 {
			return "ERR: Usage LPOP <key>"
		}
		key := parts[1]
		l := s.GetList(key)
		if l == nil {
			return "(nil)"
		}
		value, exists := l.LPop()
		if exists {
			return value
		}
		return "(nil)"

	case "RPOP":
		if len(parts) != 2 {
			return "ERR: Usage RPOP <key>"
		}
		key := parts[1]
		l := s.GetList(key)
		if l == nil {
			return "(nil)"
		}
		value, exists := l.RPop()
		if exists {
			return value
		}
		return "(nil)"

	case "SADD":
		if len(parts) < 3 {
			return "ERR: Usage SADD <key> <member> [<member>...]"
		}
		key := parts[1]
		members := parts[2:]
		set := s.GetOrCreateSet(key)
		addedCount := set.Add(members...)
		return fmt.Sprintf("%d", addedCount)

	case "SMEMBERS":
		if len(parts) != 2 {
			return "ERR: Usage SMEMBERS <key>"
		}
		key := parts[1]
		set := s.GetSet(key)
		if set == nil {
			return "(nil)"
		}
		members := set.Members()
		return strings.Join(members, " ")

	case "HSET":
		if len(parts) != 4 {
			return "ERR: Usage HSET <key> <field> <value>"
		}
		key := parts[1]
		field := parts[2]
		value := parts[3]
		hash := s.GetOrCreateHash(key)
		hash.SetField(field, value)
		return "OK"

	case "HGET":
		if len(parts) != 3 {
			return "ERR: Usage HGET <key> <field>"
		}
		key := parts[1]
		field := parts[2]
		hash := s.GetHash(key)
		if hash == nil {
			return "(nil)"
		}
		value, exists := hash.GetField(field)
		if exists {
			return value
		}
		return "(nil)"

	case "SUBSCRIBE":
		if len(parts) != 2 {
			return "ERR: Usage SUBSCRIBE <channel>"
		}
		channel := parts[1]
		ch := s.PubSub.Subscribe(channel)

		go func() {
			for msg := range ch {
				// Convert message to JSON string before sending
				jsonMsg := store.MessageToJSON(msg)
				conn.Write([]byte(fmt.Sprintf("Message %s %s\n", channel, jsonMsg)))
			}
		}()
		return "OK"

	case "PUBLISH":
		if len(parts) < 3 {
			return "ERR: Usage PUBLISH <channel> <message>"
		}
		channel := parts[1]
		message := strings.Join(parts[2:], " ")
		s.PubSub.PublishString(channel, message)
		return "OK"
	case "PUBLISH_JSON":
		if len(parts) < 3 {
			return "ERR: Usage PUBLISH_JSON <channel> <json_string>"
		}
		channel := parts[1]
		jsonStr := strings.Join(parts[2:], " ")
		var jsonData interface{}
		if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
			return "ERR: Invalid JSON"
		}
		if err := s.PubSub.PublishJSON(channel, jsonData); err != nil {
			return "ERR: " + err.Error()
		}
		return "OK"

	case "PUBLISH_INT":
		if len(parts) != 3 {
			return "ERR: Usage PUBLISH_INT <channel> <number>"
		}
		channel := parts[1]
		num, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			return "ERR: Invalid number"
		}
		s.PubSub.PublishInteger(channel, num)
		return "OK"
	case "PUBLISH_BIN":
		if len(parts) < 3 {
			return "ERR: Usage PUBLISH_BIN <channel> <data>"
		}
		channel := parts[1]
		data := []byte(strings.Join(parts[2:], " "))
		s.PubSub.PublishBinary(channel, data)
		return "OK"

	case "PUBLISH_ARRAY":
		if len(parts) < 3 {
			return "ERR: Usage PUBLISH_ARRAY <channel> <json_array>"
		}
		channel := parts[1]
		jsonArr := strings.Join(parts[2:], " ")
		var arrData []interface{}
		if err := json.Unmarshal([]byte(jsonArr), &arrData); err != nil {
			return "ERR: Invalid JSON array"
		}
		s.PubSub.PublishArray(channel, arrData)
		return "OK"
	default:
		return "ERR: Unknown command"
	}
}
