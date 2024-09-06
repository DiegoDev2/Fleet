package main

// Gotestsum - Human friendly `go test` runner
// Homepage: https://github.com/gotestyourself/gotestsum

import (
	"fmt"
	
	"os/exec"
)

func installGotestsum() {
	// Método 1: Descargar y extraer .tar.gz
	gotestsum_tar_url := "https://github.com/gotestyourself/gotestsum/archive/refs/tags/v1.12.0.tar.gz"
	gotestsum_cmd_tar := exec.Command("curl", "-L", gotestsum_tar_url, "-o", "package.tar.gz")
	err := gotestsum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gotestsum_zip_url := "https://github.com/gotestyourself/gotestsum/archive/refs/tags/v1.12.0.zip"
	gotestsum_cmd_zip := exec.Command("curl", "-L", gotestsum_zip_url, "-o", "package.zip")
	err = gotestsum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gotestsum_bin_url := "https://github.com/gotestyourself/gotestsum/archive/refs/tags/v1.12.0.bin"
	gotestsum_cmd_bin := exec.Command("curl", "-L", gotestsum_bin_url, "-o", "binary.bin")
	err = gotestsum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gotestsum_src_url := "https://github.com/gotestyourself/gotestsum/archive/refs/tags/v1.12.0.src.tar.gz"
	gotestsum_cmd_src := exec.Command("curl", "-L", gotestsum_src_url, "-o", "source.tar.gz")
	err = gotestsum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gotestsum_cmd_direct := exec.Command("./binary")
	err = gotestsum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
