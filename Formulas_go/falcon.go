package main

// Falcon - Multi-paradigm programming language and scripting engine
// Homepage: http://www.falconpl.org/

import (
	"fmt"
	
	"os/exec"
)

func installFalcon() {
	// Método 1: Descargar y extraer .tar.gz
	falcon_tar_url := "https://mirrorservice.org/sites/distfiles.macports.org/falcon/Falcon-0.9.6.8.tgz"
	falcon_cmd_tar := exec.Command("curl", "-L", falcon_tar_url, "-o", "package.tar.gz")
	err := falcon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	falcon_zip_url := "https://mirrorservice.org/sites/distfiles.macports.org/falcon/Falcon-0.9.6.8.tgz"
	falcon_cmd_zip := exec.Command("curl", "-L", falcon_zip_url, "-o", "package.zip")
	err = falcon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	falcon_bin_url := "https://mirrorservice.org/sites/distfiles.macports.org/falcon/Falcon-0.9.6.8.tgz"
	falcon_cmd_bin := exec.Command("curl", "-L", falcon_bin_url, "-o", "binary.bin")
	err = falcon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	falcon_src_url := "https://mirrorservice.org/sites/distfiles.macports.org/falcon/Falcon-0.9.6.8.tgz"
	falcon_cmd_src := exec.Command("curl", "-L", falcon_src_url, "-o", "source.tar.gz")
	err = falcon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	falcon_cmd_direct := exec.Command("./binary")
	err = falcon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: mysql@5.7")
exec.Command("latte", "install", "mysql@5.7")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
}
