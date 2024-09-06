package main

// IosClassGuard - Objective-C obfuscator for Mach-O executables
// Homepage: https://github.com/Polidea/ios-class-guard/

import (
	"fmt"
	
	"os/exec"
)

func installIosClassGuard() {
	// Método 1: Descargar y extraer .tar.gz
	iosclassguard_tar_url := "https://github.com/Polidea/ios-class-guard/archive/refs/tags/0.8.tar.gz"
	iosclassguard_cmd_tar := exec.Command("curl", "-L", iosclassguard_tar_url, "-o", "package.tar.gz")
	err := iosclassguard_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iosclassguard_zip_url := "https://github.com/Polidea/ios-class-guard/archive/refs/tags/0.8.zip"
	iosclassguard_cmd_zip := exec.Command("curl", "-L", iosclassguard_zip_url, "-o", "package.zip")
	err = iosclassguard_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iosclassguard_bin_url := "https://github.com/Polidea/ios-class-guard/archive/refs/tags/0.8.bin"
	iosclassguard_cmd_bin := exec.Command("curl", "-L", iosclassguard_bin_url, "-o", "binary.bin")
	err = iosclassguard_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iosclassguard_src_url := "https://github.com/Polidea/ios-class-guard/archive/refs/tags/0.8.src.tar.gz"
	iosclassguard_cmd_src := exec.Command("curl", "-L", iosclassguard_src_url, "-o", "source.tar.gz")
	err = iosclassguard_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iosclassguard_cmd_direct := exec.Command("./binary")
	err = iosclassguard_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
