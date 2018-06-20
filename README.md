# go-gps-filter

Disregard potentially erroneous latitude, longitude points for a courier. This application uses simple standard deviation based approach  filter outlier points. A more advanced algorithm like Kalman's Algorithm can also be used to filter data.

__Steps__
* Calculates distance between points
* Calculates speed using the distance and time diffrence
* Discards the speed with zero value
* Finds the standard deviation of speeds
* Discards any speed > 2xSD to filter the outliers


