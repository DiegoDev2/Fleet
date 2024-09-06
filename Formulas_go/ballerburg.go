package main

// Ballerburg - Castle combat game
// Homepage: https://baller.tuxfamily.org/

import (
	"fmt"
	
	"os/exec"
)

func installBallerburg() {
	// Método 1: Descargar y extraer .tar.gz
	ballerburg_tar_url := "https://download.tuxfamily.org/baller/ballerburg-1.2.2.tar.gz"
	ballerburg_cmd_tar := exec.Command("curl", "-L", ballerburg_tar_url, "-o", "package.tar.gz")
	err := ballerburg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ballerburg_zip_url := "https://download.tuxfamily.org/baller/ballerburg-1.2.2.zip"
	ballerburg_cmd_zip := exec.Command("curl", "-L", ballerburg_zip_url, "-o", "package.zip")
	err = ballerburg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ballerburg_bin_url := "https://download.tuxfamily.org/baller/ballerburg-1.2.2.bin"
	ballerburg_cmd_bin := exec.Command("curl", "-L", ballerburg_bin_url, "-o", "binary.bin")
	err = ballerburg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ballerburg_src_url := "https://download.tuxfamily.org/baller/ballerburg-1.2.2.src.tar.gz"
	ballerburg_cmd_src := exec.Command("curl", "-L", ballerburg_src_url, "-o", "source.tar.gz")
	err = ballerburg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ballerburg_cmd_direct := exec.Command("./binary")
	err = ballerburg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
