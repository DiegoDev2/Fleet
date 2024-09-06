package main

// UserspaceRcu - Library for userspace RCU (read-copy-update)
// Homepage: https://liburcu.org

import (
	"fmt"
	
	"os/exec"
)

func installUserspaceRcu() {
	// Método 1: Descargar y extraer .tar.gz
	userspacercu_tar_url := "https://lttng.org/files/urcu/userspace-rcu-0.14.1.tar.bz2"
	userspacercu_cmd_tar := exec.Command("curl", "-L", userspacercu_tar_url, "-o", "package.tar.gz")
	err := userspacercu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	userspacercu_zip_url := "https://lttng.org/files/urcu/userspace-rcu-0.14.1.tar.bz2"
	userspacercu_cmd_zip := exec.Command("curl", "-L", userspacercu_zip_url, "-o", "package.zip")
	err = userspacercu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	userspacercu_bin_url := "https://lttng.org/files/urcu/userspace-rcu-0.14.1.tar.bz2"
	userspacercu_cmd_bin := exec.Command("curl", "-L", userspacercu_bin_url, "-o", "binary.bin")
	err = userspacercu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	userspacercu_src_url := "https://lttng.org/files/urcu/userspace-rcu-0.14.1.tar.bz2"
	userspacercu_cmd_src := exec.Command("curl", "-L", userspacercu_src_url, "-o", "source.tar.gz")
	err = userspacercu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	userspacercu_cmd_direct := exec.Command("./binary")
	err = userspacercu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
