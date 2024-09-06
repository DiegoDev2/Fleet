package main

// Snapraid - Backup program for disk arrays
// Homepage: https://www.snapraid.it/

import (
	"fmt"
	
	"os/exec"
)

func installSnapraid() {
	// Método 1: Descargar y extraer .tar.gz
	snapraid_tar_url := "https://github.com/amadvance/snapraid/releases/download/v12.3/snapraid-12.3.tar.gz"
	snapraid_cmd_tar := exec.Command("curl", "-L", snapraid_tar_url, "-o", "package.tar.gz")
	err := snapraid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snapraid_zip_url := "https://github.com/amadvance/snapraid/releases/download/v12.3/snapraid-12.3.zip"
	snapraid_cmd_zip := exec.Command("curl", "-L", snapraid_zip_url, "-o", "package.zip")
	err = snapraid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snapraid_bin_url := "https://github.com/amadvance/snapraid/releases/download/v12.3/snapraid-12.3.bin"
	snapraid_cmd_bin := exec.Command("curl", "-L", snapraid_bin_url, "-o", "binary.bin")
	err = snapraid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snapraid_src_url := "https://github.com/amadvance/snapraid/releases/download/v12.3/snapraid-12.3.src.tar.gz"
	snapraid_cmd_src := exec.Command("curl", "-L", snapraid_src_url, "-o", "source.tar.gz")
	err = snapraid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snapraid_cmd_direct := exec.Command("./binary")
	err = snapraid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
