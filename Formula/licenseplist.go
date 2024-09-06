package main

// Licenseplist - License list generator of all your dependencies for iOS applications
// Homepage: https://www.slideshare.net/mono0926/licenseplist-a-license-list-generator-of-all-your-dependencies-for-ios-applications

import (
	"fmt"
	
	"os/exec"
)

func installLicenseplist() {
	// Método 1: Descargar y extraer .tar.gz
	licenseplist_tar_url := "https://github.com/mono0926/LicensePlist/archive/refs/tags/3.25.1.tar.gz"
	licenseplist_cmd_tar := exec.Command("curl", "-L", licenseplist_tar_url, "-o", "package.tar.gz")
	err := licenseplist_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	licenseplist_zip_url := "https://github.com/mono0926/LicensePlist/archive/refs/tags/3.25.1.zip"
	licenseplist_cmd_zip := exec.Command("curl", "-L", licenseplist_zip_url, "-o", "package.zip")
	err = licenseplist_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	licenseplist_bin_url := "https://github.com/mono0926/LicensePlist/archive/refs/tags/3.25.1.bin"
	licenseplist_cmd_bin := exec.Command("curl", "-L", licenseplist_bin_url, "-o", "binary.bin")
	err = licenseplist_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	licenseplist_src_url := "https://github.com/mono0926/LicensePlist/archive/refs/tags/3.25.1.src.tar.gz"
	licenseplist_cmd_src := exec.Command("curl", "-L", licenseplist_src_url, "-o", "source.tar.gz")
	err = licenseplist_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	licenseplist_cmd_direct := exec.Command("./binary")
	err = licenseplist_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
