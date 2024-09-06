package main

// Minidlna - Media server software, compliant with DLNA/UPnP-AV clients
// Homepage: https://sourceforge.net/projects/minidlna/

import (
	"fmt"
	
	"os/exec"
)

func installMinidlna() {
	// Método 1: Descargar y extraer .tar.gz
	minidlna_tar_url := "https://downloads.sourceforge.net/project/minidlna/minidlna/1.3.3/minidlna-1.3.3.tar.gz"
	minidlna_cmd_tar := exec.Command("curl", "-L", minidlna_tar_url, "-o", "package.tar.gz")
	err := minidlna_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minidlna_zip_url := "https://downloads.sourceforge.net/project/minidlna/minidlna/1.3.3/minidlna-1.3.3.zip"
	minidlna_cmd_zip := exec.Command("curl", "-L", minidlna_zip_url, "-o", "package.zip")
	err = minidlna_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minidlna_bin_url := "https://downloads.sourceforge.net/project/minidlna/minidlna/1.3.3/minidlna-1.3.3.bin"
	minidlna_cmd_bin := exec.Command("curl", "-L", minidlna_bin_url, "-o", "binary.bin")
	err = minidlna_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minidlna_src_url := "https://downloads.sourceforge.net/project/minidlna/minidlna/1.3.3/minidlna-1.3.3.src.tar.gz"
	minidlna_cmd_src := exec.Command("curl", "-L", minidlna_src_url, "-o", "source.tar.gz")
	err = minidlna_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minidlna_cmd_direct := exec.Command("./binary")
	err = minidlna_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: ffmpeg@6")
exec.Command("latte", "install", "ffmpeg@6")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libexif")
exec.Command("latte", "install", "libexif")
	fmt.Println("Instalando dependencia: libid3tag")
exec.Command("latte", "install", "libid3tag")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
