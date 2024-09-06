package main

// Duply - Frontend to the duplicity backup system
// Homepage: https://sourceforge.net/projects/ftplicity/

import (
	"fmt"
	
	"os/exec"
)

func installDuply() {
	// Método 1: Descargar y extraer .tar.gz
	duply_tar_url := "https://downloads.sourceforge.net/project/ftplicity/duply%20%28simple%20duplicity%29/2.5.x/duply_2.5.3.tgz"
	duply_cmd_tar := exec.Command("curl", "-L", duply_tar_url, "-o", "package.tar.gz")
	err := duply_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	duply_zip_url := "https://downloads.sourceforge.net/project/ftplicity/duply%20%28simple%20duplicity%29/2.5.x/duply_2.5.3.tgz"
	duply_cmd_zip := exec.Command("curl", "-L", duply_zip_url, "-o", "package.zip")
	err = duply_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	duply_bin_url := "https://downloads.sourceforge.net/project/ftplicity/duply%20%28simple%20duplicity%29/2.5.x/duply_2.5.3.tgz"
	duply_cmd_bin := exec.Command("curl", "-L", duply_bin_url, "-o", "binary.bin")
	err = duply_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	duply_src_url := "https://downloads.sourceforge.net/project/ftplicity/duply%20%28simple%20duplicity%29/2.5.x/duply_2.5.3.tgz"
	duply_cmd_src := exec.Command("curl", "-L", duply_src_url, "-o", "source.tar.gz")
	err = duply_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	duply_cmd_direct := exec.Command("./binary")
	err = duply_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: duplicity")
	exec.Command("latte", "install", "duplicity").Run()
}
