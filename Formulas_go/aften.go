package main

// Aften - Audio encoder which generates ATSC A/52 compressed audio streams
// Homepage: https://aften.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installAften() {
	// Método 1: Descargar y extraer .tar.gz
	aften_tar_url := "https://downloads.sourceforge.net/project/aften/aften/0.0.8/aften-0.0.8.tar.bz2"
	aften_cmd_tar := exec.Command("curl", "-L", aften_tar_url, "-o", "package.tar.gz")
	err := aften_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aften_zip_url := "https://downloads.sourceforge.net/project/aften/aften/0.0.8/aften-0.0.8.tar.bz2"
	aften_cmd_zip := exec.Command("curl", "-L", aften_zip_url, "-o", "package.zip")
	err = aften_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aften_bin_url := "https://downloads.sourceforge.net/project/aften/aften/0.0.8/aften-0.0.8.tar.bz2"
	aften_cmd_bin := exec.Command("curl", "-L", aften_bin_url, "-o", "binary.bin")
	err = aften_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aften_src_url := "https://downloads.sourceforge.net/project/aften/aften/0.0.8/aften-0.0.8.tar.bz2"
	aften_cmd_src := exec.Command("curl", "-L", aften_src_url, "-o", "source.tar.gz")
	err = aften_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aften_cmd_direct := exec.Command("./binary")
	err = aften_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
