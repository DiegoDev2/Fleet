package main

// Freeipmi - In-band and out-of-band IPMI (v1.5/2.0) software
// Homepage: https://www.gnu.org/software/freeipmi/

import (
	"fmt"
	
	"os/exec"
)

func installFreeipmi() {
	// Método 1: Descargar y extraer .tar.gz
	freeipmi_tar_url := "https://ftp.gnu.org/gnu/freeipmi/freeipmi-1.6.14.tar.gz"
	freeipmi_cmd_tar := exec.Command("curl", "-L", freeipmi_tar_url, "-o", "package.tar.gz")
	err := freeipmi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freeipmi_zip_url := "https://ftp.gnu.org/gnu/freeipmi/freeipmi-1.6.14.zip"
	freeipmi_cmd_zip := exec.Command("curl", "-L", freeipmi_zip_url, "-o", "package.zip")
	err = freeipmi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freeipmi_bin_url := "https://ftp.gnu.org/gnu/freeipmi/freeipmi-1.6.14.bin"
	freeipmi_cmd_bin := exec.Command("curl", "-L", freeipmi_bin_url, "-o", "binary.bin")
	err = freeipmi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freeipmi_src_url := "https://ftp.gnu.org/gnu/freeipmi/freeipmi-1.6.14.src.tar.gz"
	freeipmi_cmd_src := exec.Command("curl", "-L", freeipmi_src_url, "-o", "source.tar.gz")
	err = freeipmi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freeipmi_cmd_direct := exec.Command("./binary")
	err = freeipmi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: argp-standalone")
	exec.Command("latte", "install", "argp-standalone").Run()
}
