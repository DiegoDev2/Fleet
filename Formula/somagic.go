package main

// Somagic - Linux capture program for the Somagic variants of EasyCAP
// Homepage: https://code.google.com/archive/p/easycap-somagic-linux/

import (
	"fmt"
	
	"os/exec"
)

func installSomagic() {
	// Método 1: Descargar y extraer .tar.gz
	somagic_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/easycap-somagic-linux/somagic-easycap_1.1.tar.gz"
	somagic_cmd_tar := exec.Command("curl", "-L", somagic_tar_url, "-o", "package.tar.gz")
	err := somagic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	somagic_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/easycap-somagic-linux/somagic-easycap_1.1.zip"
	somagic_cmd_zip := exec.Command("curl", "-L", somagic_zip_url, "-o", "package.zip")
	err = somagic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	somagic_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/easycap-somagic-linux/somagic-easycap_1.1.bin"
	somagic_cmd_bin := exec.Command("curl", "-L", somagic_bin_url, "-o", "binary.bin")
	err = somagic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	somagic_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/easycap-somagic-linux/somagic-easycap_1.1.src.tar.gz"
	somagic_cmd_src := exec.Command("curl", "-L", somagic_src_url, "-o", "source.tar.gz")
	err = somagic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	somagic_cmd_direct := exec.Command("./binary")
	err = somagic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: somagic-tools")
	exec.Command("latte", "install", "somagic-tools").Run()
}
