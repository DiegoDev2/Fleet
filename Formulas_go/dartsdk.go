package main

// DartSdk - Dart Language SDK, including the VM, dart2js, core libraries, and more
// Homepage: https://dart.dev

import (
	"fmt"
	
	"os/exec"
)

func installDartSdk() {
	// Método 1: Descargar y extraer .tar.gz
	dartsdk_tar_url := "https://github.com/dart-lang/sdk/archive/refs/tags/3.5.2.tar.gz"
	dartsdk_cmd_tar := exec.Command("curl", "-L", dartsdk_tar_url, "-o", "package.tar.gz")
	err := dartsdk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dartsdk_zip_url := "https://github.com/dart-lang/sdk/archive/refs/tags/3.5.2.zip"
	dartsdk_cmd_zip := exec.Command("curl", "-L", dartsdk_zip_url, "-o", "package.zip")
	err = dartsdk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dartsdk_bin_url := "https://github.com/dart-lang/sdk/archive/refs/tags/3.5.2.bin"
	dartsdk_cmd_bin := exec.Command("curl", "-L", dartsdk_bin_url, "-o", "binary.bin")
	err = dartsdk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dartsdk_src_url := "https://github.com/dart-lang/sdk/archive/refs/tags/3.5.2.src.tar.gz"
	dartsdk_cmd_src := exec.Command("curl", "-L", dartsdk_src_url, "-o", "source.tar.gz")
	err = dartsdk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dartsdk_cmd_direct := exec.Command("./binary")
	err = dartsdk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
