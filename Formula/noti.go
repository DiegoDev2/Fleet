package main

// Noti - Trigger notifications when a process completes
// Homepage: https://github.com/variadico/noti

import (
	"fmt"
	
	"os/exec"
)

func installNoti() {
	// Método 1: Descargar y extraer .tar.gz
	noti_tar_url := "https://github.com/variadico/noti/archive/refs/tags/3.7.0.tar.gz"
	noti_cmd_tar := exec.Command("curl", "-L", noti_tar_url, "-o", "package.tar.gz")
	err := noti_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	noti_zip_url := "https://github.com/variadico/noti/archive/refs/tags/3.7.0.zip"
	noti_cmd_zip := exec.Command("curl", "-L", noti_zip_url, "-o", "package.zip")
	err = noti_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	noti_bin_url := "https://github.com/variadico/noti/archive/refs/tags/3.7.0.bin"
	noti_cmd_bin := exec.Command("curl", "-L", noti_bin_url, "-o", "binary.bin")
	err = noti_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	noti_src_url := "https://github.com/variadico/noti/archive/refs/tags/3.7.0.src.tar.gz"
	noti_cmd_src := exec.Command("curl", "-L", noti_src_url, "-o", "source.tar.gz")
	err = noti_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	noti_cmd_direct := exec.Command("./binary")
	err = noti_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
