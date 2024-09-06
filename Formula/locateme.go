package main

// Locateme - Find your location using Apple's geolocation services
// Homepage: https://iharder.sourceforge.net/current/macosx/locateme/

import (
	"fmt"
	
	"os/exec"
)

func installLocateme() {
	// Método 1: Descargar y extraer .tar.gz
	locateme_tar_url := "https://downloads.sourceforge.net/project/iharder/locateme/LocateMe-v0.2.1.zip"
	locateme_cmd_tar := exec.Command("curl", "-L", locateme_tar_url, "-o", "package.tar.gz")
	err := locateme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	locateme_zip_url := "https://downloads.sourceforge.net/project/iharder/locateme/LocateMe-v0.2.1.zip"
	locateme_cmd_zip := exec.Command("curl", "-L", locateme_zip_url, "-o", "package.zip")
	err = locateme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	locateme_bin_url := "https://downloads.sourceforge.net/project/iharder/locateme/LocateMe-v0.2.1.zip"
	locateme_cmd_bin := exec.Command("curl", "-L", locateme_bin_url, "-o", "binary.bin")
	err = locateme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	locateme_src_url := "https://downloads.sourceforge.net/project/iharder/locateme/LocateMe-v0.2.1.zip"
	locateme_cmd_src := exec.Command("curl", "-L", locateme_src_url, "-o", "source.tar.gz")
	err = locateme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	locateme_cmd_direct := exec.Command("./binary")
	err = locateme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
