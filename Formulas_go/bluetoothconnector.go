package main

// Bluetoothconnector - Connect and disconnect Bluetooth devices
// Homepage: https://github.com/lapfelix/BluetoothConnector

import (
	"fmt"
	
	"os/exec"
)

func installBluetoothconnector() {
	// Método 1: Descargar y extraer .tar.gz
	bluetoothconnector_tar_url := "https://github.com/lapfelix/BluetoothConnector/archive/refs/tags/2.1.0.tar.gz"
	bluetoothconnector_cmd_tar := exec.Command("curl", "-L", bluetoothconnector_tar_url, "-o", "package.tar.gz")
	err := bluetoothconnector_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bluetoothconnector_zip_url := "https://github.com/lapfelix/BluetoothConnector/archive/refs/tags/2.1.0.zip"
	bluetoothconnector_cmd_zip := exec.Command("curl", "-L", bluetoothconnector_zip_url, "-o", "package.zip")
	err = bluetoothconnector_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bluetoothconnector_bin_url := "https://github.com/lapfelix/BluetoothConnector/archive/refs/tags/2.1.0.bin"
	bluetoothconnector_cmd_bin := exec.Command("curl", "-L", bluetoothconnector_bin_url, "-o", "binary.bin")
	err = bluetoothconnector_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bluetoothconnector_src_url := "https://github.com/lapfelix/BluetoothConnector/archive/refs/tags/2.1.0.src.tar.gz"
	bluetoothconnector_cmd_src := exec.Command("curl", "-L", bluetoothconnector_src_url, "-o", "source.tar.gz")
	err = bluetoothconnector_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bluetoothconnector_cmd_direct := exec.Command("./binary")
	err = bluetoothconnector_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
