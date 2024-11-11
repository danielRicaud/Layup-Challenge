# Basic Flight Simulation

This project contains a simple 2D flight simulation implemented in `basicFlight.html`. The simulation calculates the trajectory and yaw angle of an airplane.

## How to Run

1. Open the `basicFlight.html` file in a web browser.
2. The simulation will automatically start and display the flight path of the airplane.

## Trajectory and Yaw Angle Logic

The trajectory and yaw angle of the airplane are calculated using basic physics formulas. Here is a breakdown of the logic:

### Variables

- `initialVelocity` (v): The speed of the airplane in knots.
- `yawAngle` (θ): The yaw angle in radians at which the airplane is flying.

### Formulas

1. **Calculating X and Y as the plane moves across the canvas:**

- `newX = velocity * cos(θ)`
- `newY = velocity * sin(θ)`

### Angle Adjustment

Yaw angle is measured relative to the north, as a yaw angle of zero degrees indicates traveling northbound, while on a cartesian plane (like an html canvas) an angle of zero degrees indicates travel to the right in an eastbound direction.

In order to correctly calculate the yaw angle it had to be ensured that the zero degrees angle is aligned with the north direction rather than the east on the canvas. So the airplane's angle had to be offset to the left by `Math.PI / 2`.

<img src="https://images.slideplayer.com/15/4747240/slides/slide_2.jpg" alt="Description" width="400" height="300">

### Example Calculation

Given:

- `initialVelocity = 50 knots`
- `yawAngle = 45°`

1. Convert angle to radians: `θ = 45 * (π / 180) = 0.785 radians`
2. Calculate horizontal and vertical components:

- `newX += 50 * cos(0.785)`
- `newY += 50 * sin(0.785)`

### Note on Units

In this simulation, the velocity is measured in knots, which is a common unit of speed used in aviation and maritime contexts. One knot is equivalent to one nautical mile per hour.

### Yaw Angle

All angle references in this simulation pertain to the yaw angle, which is the angle of rotation around the vertical axis. This angle determines the direction in which the projectile is heading, with zero degrees indicating north.
