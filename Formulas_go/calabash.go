package main

// Calabash - XProc (XML Pipeline Language) implementation
// Homepage: https://xmlcalabash.com/

import (
	"fmt"
	
	"os/exec"
)

func installCalabash() {
	// Método 1: Descargar y extraer .tar.gz
	calabash_tar_url := "https://github.com/ndw/xmlcalabash1/releases/download/1.5.7-120/xmlcalabash-1.5.7-120.zip"
	calabash_cmd_tar := exec.Command("curl", "-L", calabash_tar_url, "-o", "package.tar.gz")
	err := calabash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	calabash_zip_url := "https://github.com/ndw/xmlcalabash1/releases/download/1.5.7-120/xmlcalabash-1.5.7-120.zip"
	calabash_cmd_zip := exec.Command("curl", "-L", calabash_zip_url, "-o", "package.zip")
	err = calabash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	calabash_bin_url := "https://github.com/ndw/xmlcalabash1/releases/download/1.5.7-120/xmlcalabash-1.5.7-120.zip"
	calabash_cmd_bin := exec.Command("curl", "-L", calabash_bin_url, "-o", "binary.bin")
	err = calabash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	calabash_src_url := "https://github.com/ndw/xmlcalabash1/releases/download/1.5.7-120/xmlcalabash-1.5.7-120.zip"
	calabash_cmd_src := exec.Command("curl", "-L", calabash_src_url, "-o", "source.tar.gz")
	err = calabash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	calabash_cmd_direct := exec.Command("./binary")
	err = calabash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: saxon")
exec.Command("latte", "install", "saxon")
}
