package main

// OpenAdventure - Colossal Cave Adventure, the 1995 430-point version
// Homepage: http://www.catb.org/~esr/open-adventure/

import (
	"fmt"
	
	"os/exec"
)

func installOpenAdventure() {
	// Método 1: Descargar y extraer .tar.gz
	openadventure_tar_url := "http://www.catb.org/~esr/open-adventure/advent-1.19.tar.gz"
	openadventure_cmd_tar := exec.Command("curl", "-L", openadventure_tar_url, "-o", "package.tar.gz")
	err := openadventure_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openadventure_zip_url := "http://www.catb.org/~esr/open-adventure/advent-1.19.zip"
	openadventure_cmd_zip := exec.Command("curl", "-L", openadventure_zip_url, "-o", "package.zip")
	err = openadventure_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openadventure_bin_url := "http://www.catb.org/~esr/open-adventure/advent-1.19.bin"
	openadventure_cmd_bin := exec.Command("curl", "-L", openadventure_bin_url, "-o", "binary.bin")
	err = openadventure_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openadventure_src_url := "http://www.catb.org/~esr/open-adventure/advent-1.19.src.tar.gz"
	openadventure_cmd_src := exec.Command("curl", "-L", openadventure_src_url, "-o", "source.tar.gz")
	err = openadventure_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openadventure_cmd_direct := exec.Command("./binary")
	err = openadventure_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
