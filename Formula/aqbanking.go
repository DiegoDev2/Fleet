package main

// Aqbanking - Generic online banking interface
// Homepage: https://www.aquamaniac.de/rdm/projects/aqbanking

import (
	"fmt"
	
	"os/exec"
)

func installAqbanking() {
	// Método 1: Descargar y extraer .tar.gz
	aqbanking_tar_url := "https://www.aquamaniac.de/rdm/attachments/download/499/aqbanking-6.5.4.tar.gz"
	aqbanking_cmd_tar := exec.Command("curl", "-L", aqbanking_tar_url, "-o", "package.tar.gz")
	err := aqbanking_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aqbanking_zip_url := "https://www.aquamaniac.de/rdm/attachments/download/499/aqbanking-6.5.4.zip"
	aqbanking_cmd_zip := exec.Command("curl", "-L", aqbanking_zip_url, "-o", "package.zip")
	err = aqbanking_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aqbanking_bin_url := "https://www.aquamaniac.de/rdm/attachments/download/499/aqbanking-6.5.4.bin"
	aqbanking_cmd_bin := exec.Command("curl", "-L", aqbanking_bin_url, "-o", "binary.bin")
	err = aqbanking_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aqbanking_src_url := "https://www.aquamaniac.de/rdm/attachments/download/499/aqbanking-6.5.4.src.tar.gz"
	aqbanking_cmd_src := exec.Command("curl", "-L", aqbanking_src_url, "-o", "source.tar.gz")
	err = aqbanking_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aqbanking_cmd_direct := exec.Command("./binary")
	err = aqbanking_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: gwenhywfar")
	exec.Command("latte", "install", "gwenhywfar").Run()
	fmt.Println("Instalando dependencia: ktoblzcheck")
	exec.Command("latte", "install", "ktoblzcheck").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: libxmlsec1")
	exec.Command("latte", "install", "libxmlsec1").Run()
	fmt.Println("Instalando dependencia: libxslt")
	exec.Command("latte", "install", "libxslt").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
