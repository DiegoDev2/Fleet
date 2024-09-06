package main

// Iftop - Display an interface's bandwidth usage
// Homepage: https://pdw.ex-parrot.com/iftop/

import (
	"fmt"
	
	"os/exec"
)

func installIftop() {
	// Método 1: Descargar y extraer .tar.gz
	iftop_tar_url := "https://pdw.ex-parrot.com/iftop/download/iftop-1.0pre4.tar.gz"
	iftop_cmd_tar := exec.Command("curl", "-L", iftop_tar_url, "-o", "package.tar.gz")
	err := iftop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iftop_zip_url := "https://pdw.ex-parrot.com/iftop/download/iftop-1.0pre4.zip"
	iftop_cmd_zip := exec.Command("curl", "-L", iftop_zip_url, "-o", "package.zip")
	err = iftop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iftop_bin_url := "https://pdw.ex-parrot.com/iftop/download/iftop-1.0pre4.bin"
	iftop_cmd_bin := exec.Command("curl", "-L", iftop_bin_url, "-o", "binary.bin")
	err = iftop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iftop_src_url := "https://pdw.ex-parrot.com/iftop/download/iftop-1.0pre4.src.tar.gz"
	iftop_cmd_src := exec.Command("curl", "-L", iftop_src_url, "-o", "source.tar.gz")
	err = iftop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iftop_cmd_direct := exec.Command("./binary")
	err = iftop_cmd_direct.Run()
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
