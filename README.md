# aircto
permitted user to login using jwt 

Logged in user can only Create a Issue 


rest End points with Test 

Login Part
localhost:3000/api/user/login/ --> Method POST
@payload
{

	"name":"muthu",
	"email":"xy@gmail.com"
}
---------------------------Verfication
localhost:3000/api/user/issue/  -->Method POST
@payload

{
	"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6Ik11dGh1In0.c5SDVsnjHXUof2Ha4J88xVXJ5GBV6MlnAJVnP0q0m9A",
	"userName":"Muthu"
}


