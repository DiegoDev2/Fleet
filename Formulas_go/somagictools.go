package main

// SomagicTools - Tools to extract firmware from EasyCAP
// Homepage: https://code.google.com/archive/p/easycap-somagic-linux/

import (
	"fmt"
	
	"os/exec"
)

func installSomagicTools() {
	// Método 1: Descargar y extraer .tar.gz
	somagictools_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/easycap-somagic-linux/somagic-easycap-tools_1.1.tar.gz"
	somagictools_cmd_tar := exec.Command("curl", "-L", somagictools_tar_url, "-o", "package.tar.gz")
	err := somagictools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	somagictools_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/easycap-somagic-linux/somagic-easycap-tools_1.1.zip"
	somagictools_cmd_zip := exec.Command("curl", "-L", somagictools_zip_url, "-o", "package.zip")
	err = somagictools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	somagictools_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/easycap-somagic-linux/somagic-easycap-tools_1.1.bin"
	somagictools_cmd_bin := exec.Command("curl", "-L", somagictools_bin_url, "-o", "binary.bin")
	err = somagictools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	somagictools_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/easycap-somagic-linux/somagic-easycap-tools_1.1.src.tar.gz"
	somagictools_cmd_src := exec.Command("curl", "-L", somagictools_src_url, "-o", "source.tar.gz")
	err = somagictools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	somagictools_cmd_direct := exec.Command("./binary")
	err = somagictools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
}
