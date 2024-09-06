package main

// Dex2jar - Tools to work with Android .dex and Java .class files
// Homepage: https://github.com/pxb1988/dex2jar

import (
	"fmt"
	
	"os/exec"
)

func installDex2jar() {
	// Método 1: Descargar y extraer .tar.gz
	dex2jar_tar_url := "https://github.com/pxb1988/dex2jar/releases/download/v2.4/dex-tools-v2.4.zip"
	dex2jar_cmd_tar := exec.Command("curl", "-L", dex2jar_tar_url, "-o", "package.tar.gz")
	err := dex2jar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dex2jar_zip_url := "https://github.com/pxb1988/dex2jar/releases/download/v2.4/dex-tools-v2.4.zip"
	dex2jar_cmd_zip := exec.Command("curl", "-L", dex2jar_zip_url, "-o", "package.zip")
	err = dex2jar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dex2jar_bin_url := "https://github.com/pxb1988/dex2jar/releases/download/v2.4/dex-tools-v2.4.zip"
	dex2jar_cmd_bin := exec.Command("curl", "-L", dex2jar_bin_url, "-o", "binary.bin")
	err = dex2jar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dex2jar_src_url := "https://github.com/pxb1988/dex2jar/releases/download/v2.4/dex-tools-v2.4.zip"
	dex2jar_cmd_src := exec.Command("curl", "-L", dex2jar_src_url, "-o", "source.tar.gz")
	err = dex2jar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dex2jar_cmd_direct := exec.Command("./binary")
	err = dex2jar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
