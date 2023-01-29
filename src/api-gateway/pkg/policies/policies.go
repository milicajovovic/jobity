package policies

import "github.com/casbin/casbin/v2"

func SetupPolicies(enforcer *casbin.Enforcer) {
	// Adding roles for users in initial databases
	enforcer.AddGroupingPolicy("admin@gmail.com", "admin")
	enforcer.AddGroupingPolicy("smart@gmail.com", "admin")
	enforcer.AddGroupingPolicy("john@doe.com", "admin")
	enforcer.AddGroupingPolicy("jane@doe.com", "employee")
	enforcer.AddGroupingPolicy("foody@gmail.com", "employer")
	enforcer.AddGroupingPolicy("abank@gmail.com", "employer")
	enforcer.AddGroupingPolicy("itsoft@gmail.com", "employer")

	// Policies for admin-service

	// Policies for ad-service

	// Policies for application-service

	// Policies for employee-service
	if hasPolicy := enforcer.HasPolicy("admin", "/employees", "GET"); !hasPolicy {
		enforcer.AddPolicy("admin", "/employees", "GET")
	}
	if hasPolicy := enforcer.HasPolicy("employee", "/employees/update", "POST"); !hasPolicy {
		enforcer.AddPolicy("employee", "/employees/update", "POST")
	}

	// Policies for employer-service

	// Policies for review-service
}
