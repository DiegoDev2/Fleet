package main

// Gotop - Terminal based graphical activity monitor inspired by gtop and vtop
// Homepage: https://github.com/xxxserxxx/gotop

import (
	"fmt"
	
	"os/exec"
)

func installGotop() {
	// Método 1: Descargar y extraer .tar.gz
	gotop_tar_url := "https://github.com/xxxserxxx/gotop/archive/refs/tags/v4.2.0.tar.gz"
	gotop_cmd_tar := exec.Command("curl", "-L", gotop_tar_url, "-o", "package.tar.gz")
	err := gotop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gotop_zip_url := "https://github.com/xxxserxxx/gotop/archive/refs/tags/v4.2.0.zip"
	gotop_cmd_zip := exec.Command("curl", "-L", gotop_zip_url, "-o", "package.zip")
	err = gotop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gotop_bin_url := "https://github.com/xxxserxxx/gotop/archive/refs/tags/v4.2.0.bin"
	gotop_cmd_bin := exec.Command("curl", "-L", gotop_bin_url, "-o", "binary.bin")
	err = gotop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gotop_src_url := "https://github.com/xxxserxxx/gotop/archive/refs/tags/v4.2.0.src.tar.gz"
	gotop_cmd_src := exec.Command("curl", "-L", gotop_src_url, "-o", "source.tar.gz")
	err = gotop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gotop_cmd_direct := exec.Command("./binary")
	err = gotop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
