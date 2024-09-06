package main

// JavaServiceWrapper - Simplify the deployment, launch and monitoring of Java applications
// Homepage: https://wrapper.tanukisoftware.com/

import (
	"fmt"
	
	"os/exec"
)

func installJavaServiceWrapper() {
	// Método 1: Descargar y extraer .tar.gz
	javaservicewrapper_tar_url := "https://downloads.sourceforge.net/project/wrapper/wrapper_src/Wrapper_3.5.59_20240723/wrapper_3.5.59_src.tar.gz"
	javaservicewrapper_cmd_tar := exec.Command("curl", "-L", javaservicewrapper_tar_url, "-o", "package.tar.gz")
	err := javaservicewrapper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	javaservicewrapper_zip_url := "https://downloads.sourceforge.net/project/wrapper/wrapper_src/Wrapper_3.5.59_20240723/wrapper_3.5.59_src.zip"
	javaservicewrapper_cmd_zip := exec.Command("curl", "-L", javaservicewrapper_zip_url, "-o", "package.zip")
	err = javaservicewrapper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	javaservicewrapper_bin_url := "https://downloads.sourceforge.net/project/wrapper/wrapper_src/Wrapper_3.5.59_20240723/wrapper_3.5.59_src.bin"
	javaservicewrapper_cmd_bin := exec.Command("curl", "-L", javaservicewrapper_bin_url, "-o", "binary.bin")
	err = javaservicewrapper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	javaservicewrapper_src_url := "https://downloads.sourceforge.net/project/wrapper/wrapper_src/Wrapper_3.5.59_20240723/wrapper_3.5.59_src.src.tar.gz"
	javaservicewrapper_cmd_src := exec.Command("curl", "-L", javaservicewrapper_src_url, "-o", "source.tar.gz")
	err = javaservicewrapper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	javaservicewrapper_cmd_direct := exec.Command("./binary")
	err = javaservicewrapper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
exec.Command("latte", "install", "ant")
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
	fmt.Println("Instalando dependencia: cunit")
exec.Command("latte", "install", "cunit")
}
