package main

// Healpix - Hierarchical Equal Area isoLatitude Pixelization of a sphere
// Homepage: https://healpix.jpl.nasa.gov

import (
	"fmt"
	
	"os/exec"
)

func installHealpix() {
	// Método 1: Descargar y extraer .tar.gz
	healpix_tar_url := "https://downloads.sourceforge.net/project/healpix/Healpix_3.82/Healpix_3.82_2022Jul28.tar.gz"
	healpix_cmd_tar := exec.Command("curl", "-L", healpix_tar_url, "-o", "package.tar.gz")
	err := healpix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	healpix_zip_url := "https://downloads.sourceforge.net/project/healpix/Healpix_3.82/Healpix_3.82_2022Jul28.zip"
	healpix_cmd_zip := exec.Command("curl", "-L", healpix_zip_url, "-o", "package.zip")
	err = healpix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	healpix_bin_url := "https://downloads.sourceforge.net/project/healpix/Healpix_3.82/Healpix_3.82_2022Jul28.bin"
	healpix_cmd_bin := exec.Command("curl", "-L", healpix_bin_url, "-o", "binary.bin")
	err = healpix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	healpix_src_url := "https://downloads.sourceforge.net/project/healpix/Healpix_3.82/Healpix_3.82_2022Jul28.src.tar.gz"
	healpix_cmd_src := exec.Command("curl", "-L", healpix_src_url, "-o", "source.tar.gz")
	err = healpix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	healpix_cmd_direct := exec.Command("./binary")
	err = healpix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cfitsio")
	exec.Command("latte", "install", "cfitsio").Run()
}
