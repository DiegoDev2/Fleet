package main

// Xplanet - Create HQ wallpapers of planet Earth
// Homepage: https://xplanet.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installXplanet() {
	// Método 1: Descargar y extraer .tar.gz
	xplanet_tar_url := "https://downloads.sourceforge.net/project/xplanet/xplanet/1.3.1/xplanet-1.3.1.tar.gz"
	xplanet_cmd_tar := exec.Command("curl", "-L", xplanet_tar_url, "-o", "package.tar.gz")
	err := xplanet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xplanet_zip_url := "https://downloads.sourceforge.net/project/xplanet/xplanet/1.3.1/xplanet-1.3.1.zip"
	xplanet_cmd_zip := exec.Command("curl", "-L", xplanet_zip_url, "-o", "package.zip")
	err = xplanet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xplanet_bin_url := "https://downloads.sourceforge.net/project/xplanet/xplanet/1.3.1/xplanet-1.3.1.bin"
	xplanet_cmd_bin := exec.Command("curl", "-L", xplanet_bin_url, "-o", "binary.bin")
	err = xplanet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xplanet_src_url := "https://downloads.sourceforge.net/project/xplanet/xplanet/1.3.1/xplanet-1.3.1.src.tar.gz"
	xplanet_cmd_src := exec.Command("curl", "-L", xplanet_src_url, "-o", "source.tar.gz")
	err = xplanet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xplanet_cmd_direct := exec.Command("./binary")
	err = xplanet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: giflib")
	exec.Command("latte", "install", "giflib").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
