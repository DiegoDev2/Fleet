package main

// Geomview - Interactive 3D viewing program
// Homepage: http://www.geomview.org

import (
	"fmt"
	
	"os/exec"
)

func installGeomview() {
	// Método 1: Descargar y extraer .tar.gz
	geomview_tar_url := "https://deb.debian.org/debian/pool/main/g/geomview/geomview_1.9.5.orig.tar.gz"
	geomview_cmd_tar := exec.Command("curl", "-L", geomview_tar_url, "-o", "package.tar.gz")
	err := geomview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geomview_zip_url := "https://deb.debian.org/debian/pool/main/g/geomview/geomview_1.9.5.orig.zip"
	geomview_cmd_zip := exec.Command("curl", "-L", geomview_zip_url, "-o", "package.zip")
	err = geomview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geomview_bin_url := "https://deb.debian.org/debian/pool/main/g/geomview/geomview_1.9.5.orig.bin"
	geomview_cmd_bin := exec.Command("curl", "-L", geomview_bin_url, "-o", "binary.bin")
	err = geomview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geomview_src_url := "https://deb.debian.org/debian/pool/main/g/geomview/geomview_1.9.5.orig.src.tar.gz"
	geomview_cmd_src := exec.Command("curl", "-L", geomview_src_url, "-o", "source.tar.gz")
	err = geomview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geomview_cmd_direct := exec.Command("./binary")
	err = geomview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libice")
	exec.Command("latte", "install", "libice").Run()
	fmt.Println("Instalando dependencia: libsm")
	exec.Command("latte", "install", "libsm").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxmu")
	exec.Command("latte", "install", "libxmu").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
	fmt.Println("Instalando dependencia: openmotif")
	exec.Command("latte", "install", "openmotif").Run()
}
