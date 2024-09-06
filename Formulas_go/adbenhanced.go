package main

// AdbEnhanced - Swiss-army knife for Android testing and development
// Homepage: https://ashishb.net/tech/introducing-adb-enhanced-a-swiss-army-knife-for-android-development/

import (
	"fmt"
	
	"os/exec"
)

func installAdbEnhanced() {
	// Método 1: Descargar y extraer .tar.gz
	adbenhanced_tar_url := "https://files.pythonhosted.org/packages/8d/7d/baff371b8795aec480c607b08473502dde2dcaf4b887f22aa30c42979293/adb-enhanced-2.5.24.tar.gz"
	adbenhanced_cmd_tar := exec.Command("curl", "-L", adbenhanced_tar_url, "-o", "package.tar.gz")
	err := adbenhanced_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	adbenhanced_zip_url := "https://files.pythonhosted.org/packages/8d/7d/baff371b8795aec480c607b08473502dde2dcaf4b887f22aa30c42979293/adb-enhanced-2.5.24.zip"
	adbenhanced_cmd_zip := exec.Command("curl", "-L", adbenhanced_zip_url, "-o", "package.zip")
	err = adbenhanced_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	adbenhanced_bin_url := "https://files.pythonhosted.org/packages/8d/7d/baff371b8795aec480c607b08473502dde2dcaf4b887f22aa30c42979293/adb-enhanced-2.5.24.bin"
	adbenhanced_cmd_bin := exec.Command("curl", "-L", adbenhanced_bin_url, "-o", "binary.bin")
	err = adbenhanced_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	adbenhanced_src_url := "https://files.pythonhosted.org/packages/8d/7d/baff371b8795aec480c607b08473502dde2dcaf4b887f22aa30c42979293/adb-enhanced-2.5.24.src.tar.gz"
	adbenhanced_cmd_src := exec.Command("curl", "-L", adbenhanced_src_url, "-o", "source.tar.gz")
	err = adbenhanced_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	adbenhanced_cmd_direct := exec.Command("./binary")
	err = adbenhanced_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
