package main

// Vncsnapshot - Command-line utility for taking VNC snapshots
// Homepage: https://sourceforge.net/projects/vncsnapshot/

import (
	"fmt"
	
	"os/exec"
)

func installVncsnapshot() {
	// Método 1: Descargar y extraer .tar.gz
	vncsnapshot_tar_url := "https://downloads.sourceforge.net/project/vncsnapshot/vncsnapshot/1.2a/vncsnapshot-1.2a-src.tar.gz"
	vncsnapshot_cmd_tar := exec.Command("curl", "-L", vncsnapshot_tar_url, "-o", "package.tar.gz")
	err := vncsnapshot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vncsnapshot_zip_url := "https://downloads.sourceforge.net/project/vncsnapshot/vncsnapshot/1.2a/vncsnapshot-1.2a-src.zip"
	vncsnapshot_cmd_zip := exec.Command("curl", "-L", vncsnapshot_zip_url, "-o", "package.zip")
	err = vncsnapshot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vncsnapshot_bin_url := "https://downloads.sourceforge.net/project/vncsnapshot/vncsnapshot/1.2a/vncsnapshot-1.2a-src.bin"
	vncsnapshot_cmd_bin := exec.Command("curl", "-L", vncsnapshot_bin_url, "-o", "binary.bin")
	err = vncsnapshot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vncsnapshot_src_url := "https://downloads.sourceforge.net/project/vncsnapshot/vncsnapshot/1.2a/vncsnapshot-1.2a-src.src.tar.gz"
	vncsnapshot_cmd_src := exec.Command("curl", "-L", vncsnapshot_src_url, "-o", "source.tar.gz")
	err = vncsnapshot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vncsnapshot_cmd_direct := exec.Command("./binary")
	err = vncsnapshot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
}
