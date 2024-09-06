package main

// ProcyonDecompiler - Modern decompiler for Java 5 and beyond
// Homepage: https://github.com/mstrobel/procyon

import (
	"fmt"
	
	"os/exec"
)

func installProcyonDecompiler() {
	// Método 1: Descargar y extraer .tar.gz
	procyondecompiler_tar_url := "https://github.com/mstrobel/procyon/releases/download/v0.6.0/procyon-decompiler-0.6.0.jar"
	procyondecompiler_cmd_tar := exec.Command("curl", "-L", procyondecompiler_tar_url, "-o", "package.tar.gz")
	err := procyondecompiler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	procyondecompiler_zip_url := "https://github.com/mstrobel/procyon/releases/download/v0.6.0/procyon-decompiler-0.6.0.jar"
	procyondecompiler_cmd_zip := exec.Command("curl", "-L", procyondecompiler_zip_url, "-o", "package.zip")
	err = procyondecompiler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	procyondecompiler_bin_url := "https://github.com/mstrobel/procyon/releases/download/v0.6.0/procyon-decompiler-0.6.0.jar"
	procyondecompiler_cmd_bin := exec.Command("curl", "-L", procyondecompiler_bin_url, "-o", "binary.bin")
	err = procyondecompiler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	procyondecompiler_src_url := "https://github.com/mstrobel/procyon/releases/download/v0.6.0/procyon-decompiler-0.6.0.jar"
	procyondecompiler_cmd_src := exec.Command("curl", "-L", procyondecompiler_src_url, "-o", "source.tar.gz")
	err = procyondecompiler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	procyondecompiler_cmd_direct := exec.Command("./binary")
	err = procyondecompiler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@21")
	exec.Command("latte", "install", "openjdk@21").Run()
}
