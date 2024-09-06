package main

// Povray - Persistence Of Vision RAYtracer (POVRAY)
// Homepage: https://www.povray.org/

import (
	"fmt"
	
	"os/exec"
)

func installPovray() {
	// Método 1: Descargar y extraer .tar.gz
	povray_tar_url := "https://github.com/POV-Ray/povray/archive/refs/tags/v3.7.0.10.tar.gz"
	povray_cmd_tar := exec.Command("curl", "-L", povray_tar_url, "-o", "package.tar.gz")
	err := povray_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	povray_zip_url := "https://github.com/POV-Ray/povray/archive/refs/tags/v3.7.0.10.zip"
	povray_cmd_zip := exec.Command("curl", "-L", povray_zip_url, "-o", "package.zip")
	err = povray_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	povray_bin_url := "https://github.com/POV-Ray/povray/archive/refs/tags/v3.7.0.10.bin"
	povray_cmd_bin := exec.Command("curl", "-L", povray_bin_url, "-o", "binary.bin")
	err = povray_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	povray_src_url := "https://github.com/POV-Ray/povray/archive/refs/tags/v3.7.0.10.src.tar.gz"
	povray_cmd_src := exec.Command("curl", "-L", povray_src_url, "-o", "source.tar.gz")
	err = povray_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	povray_cmd_direct := exec.Command("./binary")
	err = povray_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: imath")
exec.Command("latte", "install", "imath")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: openexr")
exec.Command("latte", "install", "openexr")
}
