# Find The Bad Guys

This is a simple GoLang app that gets the locations (Continent, Country, City, Subdivision) from a text list of IP addresses you provide. 

It uses the [Maxmind GeoLite Database](https://dev.maxmind.com/geoip/geoip2/geolite2/) but I did not include it as I don't have the right to redistribute it. 

You can get your own GeoIP account and get the database free of charge [by clicking here](https://www.maxmind.com/en/geolite2/signup).

It also uses this [Unofficial MaxMind GeoIP2 Reader for Go](https://github.com/oschwald/geoip2-golang)

## Steps:

Start with a text file with a list of IP addresses. If you are running CentOS you can [use this script](https://gist.github.com/JeremyMorgan/94af88899785ea725a55a382f3fd209b)

Download the Geolite database. Add its location to the **geolitedb** variable

Add the location of the IP address file to the **badguysfile** variable

Run **iplookup**

It will store the following files:

continents.txt - a list of continents
countries.txt - a list of countries
cities.txt - a list of cities
subdivisions.txt - a list of subdivisions of city


They are not checked for duplicates, in case you want to do counts on each. 

If this application is useful, [let me know](https://twitter.com/JeremyCMorgan) and I'll update it and add features. 

