package main

// Yaegi - Yet another elegant Go interpreter
// Homepage: https://github.com/traefik/yaegi

import (
	"fmt"
	
	"os/exec"
)

func installYaegi() {
	// Método 1: Descargar y extraer .tar.gz
	yaegi_tar_url := "https://github.com/traefik/yaegi/archive/refs/tags/v0.16.1.tar.gz"
	yaegi_cmd_tar := exec.Command("curl", "-L", yaegi_tar_url, "-o", "package.tar.gz")
	err := yaegi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yaegi_zip_url := "https://github.com/traefik/yaegi/archive/refs/tags/v0.16.1.zip"
	yaegi_cmd_zip := exec.Command("curl", "-L", yaegi_zip_url, "-o", "package.zip")
	err = yaegi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yaegi_bin_url := "https://github.com/traefik/yaegi/archive/refs/tags/v0.16.1.bin"
	yaegi_cmd_bin := exec.Command("curl", "-L", yaegi_bin_url, "-o", "binary.bin")
	err = yaegi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yaegi_src_url := "https://github.com/traefik/yaegi/archive/refs/tags/v0.16.1.src.tar.gz"
	yaegi_cmd_src := exec.Command("curl", "-L", yaegi_src_url, "-o", "source.tar.gz")
	err = yaegi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yaegi_cmd_direct := exec.Command("./binary")
	err = yaegi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
