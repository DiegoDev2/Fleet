package main

// Dcled - Linux driver for dream cheeky USB message board
// Homepage: https://www.jeffrika.com/~malakai/dcled/index.html

import (
	"fmt"
	
	"os/exec"
)

func installDcled() {
	// Método 1: Descargar y extraer .tar.gz
	dcled_tar_url := "https://www.jeffrika.com/~malakai/dcled/dcled-2.2.tgz"
	dcled_cmd_tar := exec.Command("curl", "-L", dcled_tar_url, "-o", "package.tar.gz")
	err := dcled_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dcled_zip_url := "https://www.jeffrika.com/~malakai/dcled/dcled-2.2.tgz"
	dcled_cmd_zip := exec.Command("curl", "-L", dcled_zip_url, "-o", "package.zip")
	err = dcled_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dcled_bin_url := "https://www.jeffrika.com/~malakai/dcled/dcled-2.2.tgz"
	dcled_cmd_bin := exec.Command("curl", "-L", dcled_bin_url, "-o", "binary.bin")
	err = dcled_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dcled_src_url := "https://www.jeffrika.com/~malakai/dcled/dcled-2.2.tgz"
	dcled_cmd_src := exec.Command("curl", "-L", dcled_src_url, "-o", "source.tar.gz")
	err = dcled_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dcled_cmd_direct := exec.Command("./binary")
	err = dcled_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
}
