package main

// Whatmask - Network settings helper
// Homepage: http://www.laffeycomputer.com/whatmask.html

import (
	"fmt"
	
	"os/exec"
)

func installWhatmask() {
	// Método 1: Descargar y extraer .tar.gz
	whatmask_tar_url := "https://web.archive.org/web/20170107110521/downloads.laffeycomputer.com/current_builds/whatmask/whatmask-1.2.tar.gz"
	whatmask_cmd_tar := exec.Command("curl", "-L", whatmask_tar_url, "-o", "package.tar.gz")
	err := whatmask_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	whatmask_zip_url := "https://web.archive.org/web/20170107110521/downloads.laffeycomputer.com/current_builds/whatmask/whatmask-1.2.zip"
	whatmask_cmd_zip := exec.Command("curl", "-L", whatmask_zip_url, "-o", "package.zip")
	err = whatmask_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	whatmask_bin_url := "https://web.archive.org/web/20170107110521/downloads.laffeycomputer.com/current_builds/whatmask/whatmask-1.2.bin"
	whatmask_cmd_bin := exec.Command("curl", "-L", whatmask_bin_url, "-o", "binary.bin")
	err = whatmask_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	whatmask_src_url := "https://web.archive.org/web/20170107110521/downloads.laffeycomputer.com/current_builds/whatmask/whatmask-1.2.src.tar.gz"
	whatmask_cmd_src := exec.Command("curl", "-L", whatmask_src_url, "-o", "source.tar.gz")
	err = whatmask_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	whatmask_cmd_direct := exec.Command("./binary")
	err = whatmask_cmd_direct.Run()
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
