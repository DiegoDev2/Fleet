package main

// WatchSim - Command-line WatchKit application launcher
// Homepage: https://github.com/alloy/watch-sim

import (
	"fmt"
	
	"os/exec"
)

func installWatchSim() {
	// Método 1: Descargar y extraer .tar.gz
	watchsim_tar_url := "https://github.com/alloy/watch-sim/archive/refs/tags/1.0.0.tar.gz"
	watchsim_cmd_tar := exec.Command("curl", "-L", watchsim_tar_url, "-o", "package.tar.gz")
	err := watchsim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	watchsim_zip_url := "https://github.com/alloy/watch-sim/archive/refs/tags/1.0.0.zip"
	watchsim_cmd_zip := exec.Command("curl", "-L", watchsim_zip_url, "-o", "package.zip")
	err = watchsim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	watchsim_bin_url := "https://github.com/alloy/watch-sim/archive/refs/tags/1.0.0.bin"
	watchsim_cmd_bin := exec.Command("curl", "-L", watchsim_bin_url, "-o", "binary.bin")
	err = watchsim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	watchsim_src_url := "https://github.com/alloy/watch-sim/archive/refs/tags/1.0.0.src.tar.gz"
	watchsim_cmd_src := exec.Command("curl", "-L", watchsim_src_url, "-o", "source.tar.gz")
	err = watchsim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	watchsim_cmd_direct := exec.Command("./binary")
	err = watchsim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
