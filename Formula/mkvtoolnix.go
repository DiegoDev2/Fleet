package main

// Mkvtoolnix - Matroska media files manipulation tools
// Homepage: https://mkvtoolnix.download/

import (
	"fmt"
	
	"os/exec"
)

func installMkvtoolnix() {
	// Método 1: Descargar y extraer .tar.gz
	mkvtoolnix_tar_url := "https://mkvtoolnix.download/sources/mkvtoolnix-86.0.tar.xz"
	mkvtoolnix_cmd_tar := exec.Command("curl", "-L", mkvtoolnix_tar_url, "-o", "package.tar.gz")
	err := mkvtoolnix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mkvtoolnix_zip_url := "https://mkvtoolnix.download/sources/mkvtoolnix-86.0.tar.xz"
	mkvtoolnix_cmd_zip := exec.Command("curl", "-L", mkvtoolnix_zip_url, "-o", "package.zip")
	err = mkvtoolnix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mkvtoolnix_bin_url := "https://mkvtoolnix.download/sources/mkvtoolnix-86.0.tar.xz"
	mkvtoolnix_cmd_bin := exec.Command("curl", "-L", mkvtoolnix_bin_url, "-o", "binary.bin")
	err = mkvtoolnix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mkvtoolnix_src_url := "https://mkvtoolnix.download/sources/mkvtoolnix-86.0.tar.xz"
	mkvtoolnix_cmd_src := exec.Command("curl", "-L", mkvtoolnix_src_url, "-o", "source.tar.gz")
	err = mkvtoolnix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mkvtoolnix_cmd_direct := exec.Command("./binary")
	err = mkvtoolnix_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: docbook-xsl")
	exec.Command("latte", "install", "docbook-xsl").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: libebml")
	exec.Command("latte", "install", "libebml").Run()
	fmt.Println("Instalando dependencia: libmatroska")
	exec.Command("latte", "install", "libmatroska").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: nlohmann-json")
	exec.Command("latte", "install", "nlohmann-json").Run()
	fmt.Println("Instalando dependencia: pugixml")
	exec.Command("latte", "install", "pugixml").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: utf8cpp")
	exec.Command("latte", "install", "utf8cpp").Run()
}
