package main

// Xsane - Graphical scanning frontend
// Homepage: https://gitlab.com/sane-project/frontend/xsane

import (
	"fmt"
	
	"os/exec"
)

func installXsane() {
	// Método 1: Descargar y extraer .tar.gz
	xsane_tar_url := "https://ftp.osuosl.org/pub/blfs/conglomeration/xsane/xsane-0.999.tar.gz"
	xsane_cmd_tar := exec.Command("curl", "-L", xsane_tar_url, "-o", "package.tar.gz")
	err := xsane_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xsane_zip_url := "https://ftp.osuosl.org/pub/blfs/conglomeration/xsane/xsane-0.999.zip"
	xsane_cmd_zip := exec.Command("curl", "-L", xsane_zip_url, "-o", "package.zip")
	err = xsane_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xsane_bin_url := "https://ftp.osuosl.org/pub/blfs/conglomeration/xsane/xsane-0.999.bin"
	xsane_cmd_bin := exec.Command("curl", "-L", xsane_bin_url, "-o", "binary.bin")
	err = xsane_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xsane_src_url := "https://ftp.osuosl.org/pub/blfs/conglomeration/xsane/xsane-0.999.src.tar.gz"
	xsane_cmd_src := exec.Command("curl", "-L", xsane_src_url, "-o", "source.tar.gz")
	err = xsane_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xsane_cmd_direct := exec.Command("./binary")
	err = xsane_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gtk+")
	exec.Command("latte", "install", "gtk+").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: sane-backends")
	exec.Command("latte", "install", "sane-backends").Run()
}
