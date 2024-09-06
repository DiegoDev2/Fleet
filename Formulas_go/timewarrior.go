package main

// Timewarrior - Command-line time tracking application
// Homepage: https://timewarrior.net/

import (
	"fmt"
	
	"os/exec"
)

func installTimewarrior() {
	// Método 1: Descargar y extraer .tar.gz
	timewarrior_tar_url := "https://github.com/GothenburgBitFactory/timewarrior/releases/download/v1.7.1/timew-1.7.1.tar.gz"
	timewarrior_cmd_tar := exec.Command("curl", "-L", timewarrior_tar_url, "-o", "package.tar.gz")
	err := timewarrior_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	timewarrior_zip_url := "https://github.com/GothenburgBitFactory/timewarrior/releases/download/v1.7.1/timew-1.7.1.zip"
	timewarrior_cmd_zip := exec.Command("curl", "-L", timewarrior_zip_url, "-o", "package.zip")
	err = timewarrior_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	timewarrior_bin_url := "https://github.com/GothenburgBitFactory/timewarrior/releases/download/v1.7.1/timew-1.7.1.bin"
	timewarrior_cmd_bin := exec.Command("curl", "-L", timewarrior_bin_url, "-o", "binary.bin")
	err = timewarrior_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	timewarrior_src_url := "https://github.com/GothenburgBitFactory/timewarrior/releases/download/v1.7.1/timew-1.7.1.src.tar.gz"
	timewarrior_cmd_src := exec.Command("curl", "-L", timewarrior_src_url, "-o", "source.tar.gz")
	err = timewarrior_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	timewarrior_cmd_direct := exec.Command("./binary")
	err = timewarrior_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoctor")
exec.Command("latte", "install", "asciidoctor")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: man-db")
exec.Command("latte", "install", "man-db")
}
