package policies

import "github.com/casbin/casbin/v2"

func SetupPolicies(enforcer *casbin.Enforcer) {
	// Adding roles for users in initial databases
	enforcer.AddGroupingPolicy("admin@gmail.com", "admin")
	enforcer.AddGroupingPolicy("smart@gmail.com", "admin")
	enforcer.AddGroupingPolicy("john@doe.com", "employee")
	enforcer.AddGroupingPolicy("jane@doe.com", "employee")
	enforcer.AddGroupingPolicy("foody@gmail.com", "employer")
	enforcer.AddGroupingPolicy("abank@gmail.com", "employer")
	enforcer.AddGroupingPolicy("itsoft@gmail.com", "employer")

	// Policies for admin-service
	enforcer.AddPolicy("admin", "/admins", "GET")
	enforcer.AddPolicy("admin", "/admins/admin/:id", "GET")

	// Policies for ad-service
	enforcer.AddPolicy("admin", "/ads/ad/:id", "GET")
	enforcer.AddPolicy("employee", "/ads/ad/:id", "GET")
	enforcer.AddPolicy("employer", "/ads/ad/:id", "GET")
	enforcer.AddPolicy("admin", "/ads/delete/:id", "POST")

	// Policies for application-service
	enforcer.AddPolicy("admin", "/applications", "GET")
	enforcer.AddPolicy("employer", "/applications/application/:id", "GET")
	enforcer.AddPolicy("employee", "/applications/apply/:adId/:employeeId", "POST")
	enforcer.AddPolicy("employee", "/applications/accepted/:id", "GET")

	// Policies for employee-service
	enforcer.AddPolicy("admin", "/employees", "GET")
	enforcer.AddPolicy("employer", "/employees/employee/:id", "GET")
	enforcer.AddPolicy("employee", "/employees/employee/:id", "GET")
	enforcer.AddPolicy("employee", "/employees/update/form", "POST")
	enforcer.AddPolicy("employee", "/employees/update/pdf", "POST")
	enforcer.AddPolicy("admin", "/employees/block/:id", "POST")
	enforcer.AddPolicy("admin", "/employees/delete/:id", "POST")

	// Policies for employer-service
	enforcer.AddPolicy("admin", "/employers", "GET")
	enforcer.AddPolicy("employee", "/employers", "GET")
	enforcer.AddPolicy("admin", "/employers/delete/:id", "POST")

	// Policies for review-service
	enforcer.AddPolicy("admin", "/reviews", "GET")
	enforcer.AddPolicy("employee", "/reviews/create", "POST")
	enforcer.AddPolicy("admin", "/reviews/appropriate/:id", "POST")
	enforcer.AddPolicy("admin", "/reviews/delete/:id", "POST")
}
