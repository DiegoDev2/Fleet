package main

// Fx - Terminal JSON viewer
// Homepage: https://fx.wtf

import (
	"fmt"
	
	"os/exec"
)

func installFx() {
	// Método 1: Descargar y extraer .tar.gz
	fx_tar_url := "https://github.com/antonmedv/fx/archive/refs/tags/35.0.0.tar.gz"
	fx_cmd_tar := exec.Command("curl", "-L", fx_tar_url, "-o", "package.tar.gz")
	err := fx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fx_zip_url := "https://github.com/antonmedv/fx/archive/refs/tags/35.0.0.zip"
	fx_cmd_zip := exec.Command("curl", "-L", fx_zip_url, "-o", "package.zip")
	err = fx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fx_bin_url := "https://github.com/antonmedv/fx/archive/refs/tags/35.0.0.bin"
	fx_cmd_bin := exec.Command("curl", "-L", fx_bin_url, "-o", "binary.bin")
	err = fx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fx_src_url := "https://github.com/antonmedv/fx/archive/refs/tags/35.0.0.src.tar.gz"
	fx_cmd_src := exec.Command("curl", "-L", fx_src_url, "-o", "source.tar.gz")
	err = fx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fx_cmd_direct := exec.Command("./binary")
	err = fx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
