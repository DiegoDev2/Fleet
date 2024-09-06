package main

// Selene - Blazing-fast modern Lua linter
// Homepage: https://kampfkarren.github.io/selene

import (
	"fmt"
	
	"os/exec"
)

func installSelene() {
	// Método 1: Descargar y extraer .tar.gz
	selene_tar_url := "https://github.com/Kampfkarren/selene/archive/refs/tags/0.27.1.tar.gz"
	selene_cmd_tar := exec.Command("curl", "-L", selene_tar_url, "-o", "package.tar.gz")
	err := selene_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	selene_zip_url := "https://github.com/Kampfkarren/selene/archive/refs/tags/0.27.1.zip"
	selene_cmd_zip := exec.Command("curl", "-L", selene_zip_url, "-o", "package.zip")
	err = selene_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	selene_bin_url := "https://github.com/Kampfkarren/selene/archive/refs/tags/0.27.1.bin"
	selene_cmd_bin := exec.Command("curl", "-L", selene_bin_url, "-o", "binary.bin")
	err = selene_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	selene_src_url := "https://github.com/Kampfkarren/selene/archive/refs/tags/0.27.1.src.tar.gz"
	selene_cmd_src := exec.Command("curl", "-L", selene_src_url, "-o", "source.tar.gz")
	err = selene_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	selene_cmd_direct := exec.Command("./binary")
	err = selene_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
