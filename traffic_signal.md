# Designing a Traffic Signal Control System

## Requirements
1. The traffic signal system should control the flow of traffic at an intersection with multiple roads.
2. The system should support different types of signals, such as red, yellow, and green.
3. The duration of each signal should be configurable and adjustable based on traffic conditions.
4. The system should handle the transition between signals smoothly, ensuring safe and efficient traffic flow.
5. The system should be able to detect and handle emergency situations, such as an ambulance or fire truck approaching the intersection.
6. The system should be scalable and extensible to support additional features and functionality.



how does the flow looks like? 

Here we are talking about only one junction? or a junction where we have 4 traffic lights and I need to manage for all 4? 

I think we are only talking about one side here? 
And in that intersection, we are going from one direction. 
That signal has some lights, what is the current light value. 
what are times set for each lights. 
If emergency situation has come, how things are going to change? 


so what all entities you think are required? 


Signal(red, yellow, green) enums

TrafficLight
    - ID
    - currentColor
    - countDown
    - ChangeSignal

Road
    - name(id) of the road
    - *TrafficLight

TrafficController
    - only exposes functions to change the colors. 
    - also has the roads array, basically which all roads it controls. 
    - And can change the signals for those

Demo? how it should look like? 

TrafficeControllerDemo
    - 5 roads it controls. 
    - road1, road2, road3, road4, road5
    - Each road signals we can change. 
    - change. 
    - isEmergency(instruction) -> (road1, finalSignal)
    



