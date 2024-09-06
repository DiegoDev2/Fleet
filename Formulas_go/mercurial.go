package main

// Mercurial - Scalable distributed version control system
// Homepage: https://mercurial-scm.org/

import (
	"fmt"
	
	"os/exec"
)

func installMercurial() {
	// Método 1: Descargar y extraer .tar.gz
	mercurial_tar_url := "https://www.mercurial-scm.org/release/mercurial-6.8.1.tar.gz"
	mercurial_cmd_tar := exec.Command("curl", "-L", mercurial_tar_url, "-o", "package.tar.gz")
	err := mercurial_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mercurial_zip_url := "https://www.mercurial-scm.org/release/mercurial-6.8.1.zip"
	mercurial_cmd_zip := exec.Command("curl", "-L", mercurial_zip_url, "-o", "package.zip")
	err = mercurial_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mercurial_bin_url := "https://www.mercurial-scm.org/release/mercurial-6.8.1.bin"
	mercurial_cmd_bin := exec.Command("curl", "-L", mercurial_bin_url, "-o", "binary.bin")
	err = mercurial_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mercurial_src_url := "https://www.mercurial-scm.org/release/mercurial-6.8.1.src.tar.gz"
	mercurial_cmd_src := exec.Command("curl", "-L", mercurial_src_url, "-o", "source.tar.gz")
	err = mercurial_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mercurial_cmd_direct := exec.Command("./binary")
	err = mercurial_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
