package main

// Shellharden - Bash syntax highlighter that encourages/fixes variables quoting
// Homepage: https://github.com/anordal/shellharden

import (
	"fmt"
	
	"os/exec"
)

func installShellharden() {
	// Método 1: Descargar y extraer .tar.gz
	shellharden_tar_url := "https://github.com/anordal/shellharden/archive/refs/tags/v4.3.1.tar.gz"
	shellharden_cmd_tar := exec.Command("curl", "-L", shellharden_tar_url, "-o", "package.tar.gz")
	err := shellharden_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shellharden_zip_url := "https://github.com/anordal/shellharden/archive/refs/tags/v4.3.1.zip"
	shellharden_cmd_zip := exec.Command("curl", "-L", shellharden_zip_url, "-o", "package.zip")
	err = shellharden_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shellharden_bin_url := "https://github.com/anordal/shellharden/archive/refs/tags/v4.3.1.bin"
	shellharden_cmd_bin := exec.Command("curl", "-L", shellharden_bin_url, "-o", "binary.bin")
	err = shellharden_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shellharden_src_url := "https://github.com/anordal/shellharden/archive/refs/tags/v4.3.1.src.tar.gz"
	shellharden_cmd_src := exec.Command("curl", "-L", shellharden_src_url, "-o", "source.tar.gz")
	err = shellharden_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shellharden_cmd_direct := exec.Command("./binary")
	err = shellharden_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
