package main

// Nb - Command-line and local web note-taking, bookmarking, and archiving
// Homepage: https://xwmx.github.io/nb

import (
	"fmt"
	
	"os/exec"
)

func installNb() {
	// Método 1: Descargar y extraer .tar.gz
	nb_tar_url := "https://github.com/xwmx/nb/archive/refs/tags/7.14.0.tar.gz"
	nb_cmd_tar := exec.Command("curl", "-L", nb_tar_url, "-o", "package.tar.gz")
	err := nb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nb_zip_url := "https://github.com/xwmx/nb/archive/refs/tags/7.14.0.zip"
	nb_cmd_zip := exec.Command("curl", "-L", nb_zip_url, "-o", "package.zip")
	err = nb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nb_bin_url := "https://github.com/xwmx/nb/archive/refs/tags/7.14.0.bin"
	nb_cmd_bin := exec.Command("curl", "-L", nb_bin_url, "-o", "binary.bin")
	err = nb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nb_src_url := "https://github.com/xwmx/nb/archive/refs/tags/7.14.0.src.tar.gz"
	nb_cmd_src := exec.Command("curl", "-L", nb_src_url, "-o", "source.tar.gz")
	err = nb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nb_cmd_direct := exec.Command("./binary")
	err = nb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bat")
	exec.Command("latte", "install", "bat").Run()
	fmt.Println("Instalando dependencia: nmap")
	exec.Command("latte", "install", "nmap").Run()
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
	fmt.Println("Instalando dependencia: ripgrep")
	exec.Command("latte", "install", "ripgrep").Run()
	fmt.Println("Instalando dependencia: tig")
	exec.Command("latte", "install", "tig").Run()
	fmt.Println("Instalando dependencia: w3m")
	exec.Command("latte", "install", "w3m").Run()
}
