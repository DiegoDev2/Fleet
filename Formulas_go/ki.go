package main

// Ki - Kotlin Language Interactive Shell
// Homepage: https://github.com/Kotlin/kotlin-interactive-shell

import (
	"fmt"
	
	"os/exec"
)

func installKi() {
	// Método 1: Descargar y extraer .tar.gz
	ki_tar_url := "https://github.com/Kotlin/kotlin-interactive-shell/archive/refs/tags/v0.5.2.tar.gz"
	ki_cmd_tar := exec.Command("curl", "-L", ki_tar_url, "-o", "package.tar.gz")
	err := ki_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ki_zip_url := "https://github.com/Kotlin/kotlin-interactive-shell/archive/refs/tags/v0.5.2.zip"
	ki_cmd_zip := exec.Command("curl", "-L", ki_zip_url, "-o", "package.zip")
	err = ki_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ki_bin_url := "https://github.com/Kotlin/kotlin-interactive-shell/archive/refs/tags/v0.5.2.bin"
	ki_cmd_bin := exec.Command("curl", "-L", ki_bin_url, "-o", "binary.bin")
	err = ki_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ki_src_url := "https://github.com/Kotlin/kotlin-interactive-shell/archive/refs/tags/v0.5.2.src.tar.gz"
	ki_cmd_src := exec.Command("curl", "-L", ki_src_url, "-o", "source.tar.gz")
	err = ki_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ki_cmd_direct := exec.Command("./binary")
	err = ki_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
exec.Command("latte", "install", "maven")
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}
