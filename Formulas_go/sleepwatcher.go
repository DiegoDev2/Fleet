package main

// Sleepwatcher - Monitors sleep, wakeup, and idleness of a Mac
// Homepage: https://www.bernhard-baehr.de/

import (
	"fmt"
	
	"os/exec"
)

func installSleepwatcher() {
	// Método 1: Descargar y extraer .tar.gz
	sleepwatcher_tar_url := "https://www.bernhard-baehr.de/sleepwatcher_2.2.1.tgz"
	sleepwatcher_cmd_tar := exec.Command("curl", "-L", sleepwatcher_tar_url, "-o", "package.tar.gz")
	err := sleepwatcher_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sleepwatcher_zip_url := "https://www.bernhard-baehr.de/sleepwatcher_2.2.1.tgz"
	sleepwatcher_cmd_zip := exec.Command("curl", "-L", sleepwatcher_zip_url, "-o", "package.zip")
	err = sleepwatcher_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sleepwatcher_bin_url := "https://www.bernhard-baehr.de/sleepwatcher_2.2.1.tgz"
	sleepwatcher_cmd_bin := exec.Command("curl", "-L", sleepwatcher_bin_url, "-o", "binary.bin")
	err = sleepwatcher_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sleepwatcher_src_url := "https://www.bernhard-baehr.de/sleepwatcher_2.2.1.tgz"
	sleepwatcher_cmd_src := exec.Command("curl", "-L", sleepwatcher_src_url, "-o", "source.tar.gz")
	err = sleepwatcher_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sleepwatcher_cmd_direct := exec.Command("./binary")
	err = sleepwatcher_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
