package main

// Dash - POSIX-compliant descendant of NetBSD's ash (the Almquist SHell)
// Homepage: http://gondor.apana.org.au/~herbert/dash/

import (
	"fmt"
	
	"os/exec"
)

func installDash() {
	// Método 1: Descargar y extraer .tar.gz
	dash_tar_url := "http://gondor.apana.org.au/~herbert/dash/files/dash-0.5.12.tar.gz"
	dash_cmd_tar := exec.Command("curl", "-L", dash_tar_url, "-o", "package.tar.gz")
	err := dash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dash_zip_url := "http://gondor.apana.org.au/~herbert/dash/files/dash-0.5.12.zip"
	dash_cmd_zip := exec.Command("curl", "-L", dash_zip_url, "-o", "package.zip")
	err = dash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dash_bin_url := "http://gondor.apana.org.au/~herbert/dash/files/dash-0.5.12.bin"
	dash_cmd_bin := exec.Command("curl", "-L", dash_bin_url, "-o", "binary.bin")
	err = dash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dash_src_url := "http://gondor.apana.org.au/~herbert/dash/files/dash-0.5.12.src.tar.gz"
	dash_cmd_src := exec.Command("curl", "-L", dash_src_url, "-o", "source.tar.gz")
	err = dash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dash_cmd_direct := exec.Command("./binary")
	err = dash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
