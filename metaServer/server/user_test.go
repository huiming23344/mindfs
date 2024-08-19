package server

import (
	"github.com/huiming23344/mindfs/metaServer/meta"
	"testing"
)

func TestAddUser_Success(t *testing.T) {
	MetaServer.Users = make(map[string]*meta.User)
	err := AddUser("user1", "password1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if _, exists := MetaServer.Users["user1"]; !exists {
		t.Fatalf("expected user 'user1' to be added")
	}
}

func TestAddUser_AlreadyExists(t *testing.T) {
	MetaServer.Users = make(map[string]*meta.User)
	_ = AddUser("user1", "password1")
	err := AddUser("user1", "password1")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestDeleteUser_Success(t *testing.T) {
	MetaServer.Users = make(map[string]*meta.User)
	_ = AddUser("user1", "password1")
	err := DeleteUser("user1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if _, exists := MetaServer.Users["user1"]; exists {
		t.Fatalf("expected user 'user1' to be deleted")
	}
}

func TestDeleteUser_NotExists(t *testing.T) {
	MetaServer.Users = make(map[string]*meta.User)
	err := DeleteUser("user1")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestListUsers(t *testing.T) {
	MetaServer.Users = make(map[string]*meta.User)
	_ = AddUser("user1", "password1")
	_ = AddUser("user2", "password2")
	users := ListUsers()
	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
}

func TestAddGroup_Success(t *testing.T) {
	MetaServer.Groups = make(map[string]*meta.UserGroup)
	err := AddGroup("group1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if _, exists := MetaServer.Groups["group1"]; !exists {
		t.Fatalf("expected group 'group1' to be added")
	}
}

func TestAddGroup_AlreadyExists(t *testing.T) {
	MetaServer.Groups = make(map[string]*meta.UserGroup)
	_ = AddGroup("group1")
	err := AddGroup("group1")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestDeleteGroup_Success(t *testing.T) {
	MetaServer.Groups = make(map[string]*meta.UserGroup)
	_ = AddGroup("group1")
	err := DeleteGroup("group1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if _, exists := MetaServer.Groups["group1"]; exists {
		t.Fatalf("expected group 'group1' to be deleted")
	}
}

func TestDeleteGroup_NotExists(t *testing.T) {
	MetaServer.Groups = make(map[string]*meta.UserGroup)
	err := DeleteGroup("group1")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestAddUserToGroup_Success(t *testing.T) {
	MetaServer.Users = make(map[string]*meta.User)
	MetaServer.Groups = make(map[string]*meta.UserGroup)
	_ = AddUser("user1", "password1")
	_ = AddGroup("group1")
	err := AddUserToGroup("user1", "group1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(MetaServer.Groups["group1"].Users) != 1 {
		t.Fatalf("expected 1 user in group 'group1', got %d", len(MetaServer.Groups["group1"].Users))
	}
}

func TestAddUserToGroup_UserNotExists(t *testing.T) {
	MetaServer.Groups = make(map[string]*meta.UserGroup)
	_ = AddGroup("group1")
	err := AddUserToGroup("user1", "group1")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestAddUserToGroup_GroupNotExists(t *testing.T) {
	MetaServer.Users = make(map[string]*meta.User)
	_ = AddUser("user1", "password1")
	err := AddUserToGroup("user1", "group1")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestRemoveUserFromGroup_Success(t *testing.T) {
	MetaServer.Users = make(map[string]*meta.User)
	MetaServer.Groups = make(map[string]*meta.UserGroup)
	_ = AddUser("user1", "password1")
	_ = AddGroup("group1")
	_ = AddUserToGroup("user1", "group1")
	err := RemoveUserFromGroup("user1", "group1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(MetaServer.Groups["group1"].Users) != 0 {
		t.Fatalf("expected 0 users in group 'group1', got %d", len(MetaServer.Groups["group1"].Users))
	}
}

func TestRemoveUserFromGroup_UserNotInGroup(t *testing.T) {
	MetaServer.Users = make(map[string]*meta.User)
	MetaServer.Groups = make(map[string]*meta.UserGroup)
	_ = AddUser("user1", "password1")
	_ = AddGroup("group1")
	err := RemoveUserFromGroup("user1", "group1")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestRemoveUserFromGroup_GroupNotExists(t *testing.T) {
	MetaServer.Users = make(map[string]*meta.User)
	_ = AddUser("user1", "password1")
	err := RemoveUserFromGroup("user1", "group1")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
