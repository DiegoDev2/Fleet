package main

// Asciiquarium - Aquarium animation in ASCII art
// Homepage: https://robobunny.com/projects/asciiquarium/html/

import (
	"fmt"
	
	"os/exec"
)

func installAsciiquarium() {
	// Método 1: Descargar y extraer .tar.gz
	asciiquarium_tar_url := "https://robobunny.com/projects/asciiquarium/asciiquarium_1.1.tar.gz"
	asciiquarium_cmd_tar := exec.Command("curl", "-L", asciiquarium_tar_url, "-o", "package.tar.gz")
	err := asciiquarium_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asciiquarium_zip_url := "https://robobunny.com/projects/asciiquarium/asciiquarium_1.1.zip"
	asciiquarium_cmd_zip := exec.Command("curl", "-L", asciiquarium_zip_url, "-o", "package.zip")
	err = asciiquarium_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asciiquarium_bin_url := "https://robobunny.com/projects/asciiquarium/asciiquarium_1.1.bin"
	asciiquarium_cmd_bin := exec.Command("curl", "-L", asciiquarium_bin_url, "-o", "binary.bin")
	err = asciiquarium_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asciiquarium_src_url := "https://robobunny.com/projects/asciiquarium/asciiquarium_1.1.src.tar.gz"
	asciiquarium_cmd_src := exec.Command("curl", "-L", asciiquarium_src_url, "-o", "source.tar.gz")
	err = asciiquarium_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asciiquarium_cmd_direct := exec.Command("./binary")
	err = asciiquarium_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: perl")
	exec.Command("latte", "install", "perl").Run()
}
