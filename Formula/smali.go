package main

// Smali - Assembler/disassembler for Android's Java VM implementation
// Homepage: https://github.com/JesusFreke/smali

import (
	"fmt"
	
	"os/exec"
)

func installSmali() {
	// Método 1: Descargar y extraer .tar.gz
	smali_tar_url := "https://github.com/JesusFreke/smali/archive/refs/tags/v2.5.2.tar.gz"
	smali_cmd_tar := exec.Command("curl", "-L", smali_tar_url, "-o", "package.tar.gz")
	err := smali_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	smali_zip_url := "https://github.com/JesusFreke/smali/archive/refs/tags/v2.5.2.zip"
	smali_cmd_zip := exec.Command("curl", "-L", smali_zip_url, "-o", "package.zip")
	err = smali_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	smali_bin_url := "https://github.com/JesusFreke/smali/archive/refs/tags/v2.5.2.bin"
	smali_cmd_bin := exec.Command("curl", "-L", smali_bin_url, "-o", "binary.bin")
	err = smali_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	smali_src_url := "https://github.com/JesusFreke/smali/archive/refs/tags/v2.5.2.src.tar.gz"
	smali_cmd_src := exec.Command("curl", "-L", smali_src_url, "-o", "source.tar.gz")
	err = smali_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	smali_cmd_direct := exec.Command("./binary")
	err = smali_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gradle@6")
	exec.Command("latte", "install", "gradle@6").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
