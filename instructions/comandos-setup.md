A. onos -> 
sudo docker exec -it <id-contenedor> ./bin/ onos -l karaf (passw: karaf)
app activate proxyarp
app activate drivers
app activate hostprovider
app activate openflow
app activate fwd

comandos: devices //  flows // hosts // nodes // paths
    
B. switch -> 
    sudo ovs-vsctl add-br brtun
    sudo ovs-vsctl set bridge brtun protocols=OpenFlow13
    sudo ovs-vsctl set-controller brtun tcp:192.168.2.10:6633
    sudo ovs-vsctl add-port brtun eth0
    sudo ovs-vsctl add-port brtun eth1

    auxliar:

    sudo ovs-vsctl show -> Muestra el bridge y si está conectado 
    sudo ovs-ofctl dump-flows brtun --protocol=OpenFlow13 --> Mostrar las reglas del switch


C. QKADS ->sudo docker exec -it 

    en uno de los dos hay que poner una ip de la red del otro: 
        sudo ip a add 192.168.0.20/24 dev eth0