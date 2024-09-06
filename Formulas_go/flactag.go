package main

// Flactag - Tag single album FLAC files with MusicBrainz CUE sheets
// Homepage: https://flactag.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installFlactag() {
	// Método 1: Descargar y extraer .tar.gz
	flactag_tar_url := "https://downloads.sourceforge.net/project/flactag/v2.0.4/flactag-2.0.4.tar.gz"
	flactag_cmd_tar := exec.Command("curl", "-L", flactag_tar_url, "-o", "package.tar.gz")
	err := flactag_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flactag_zip_url := "https://downloads.sourceforge.net/project/flactag/v2.0.4/flactag-2.0.4.zip"
	flactag_cmd_zip := exec.Command("curl", "-L", flactag_zip_url, "-o", "package.zip")
	err = flactag_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flactag_bin_url := "https://downloads.sourceforge.net/project/flactag/v2.0.4/flactag-2.0.4.bin"
	flactag_cmd_bin := exec.Command("curl", "-L", flactag_bin_url, "-o", "binary.bin")
	err = flactag_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flactag_src_url := "https://downloads.sourceforge.net/project/flactag/v2.0.4/flactag-2.0.4.src.tar.gz"
	flactag_cmd_src := exec.Command("curl", "-L", flactag_src_url, "-o", "source.tar.gz")
	err = flactag_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flactag_cmd_direct := exec.Command("./binary")
	err = flactag_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libdiscid")
exec.Command("latte", "install", "libdiscid")
	fmt.Println("Instalando dependencia: libmusicbrainz")
exec.Command("latte", "install", "libmusicbrainz")
	fmt.Println("Instalando dependencia: neon")
exec.Command("latte", "install", "neon")
	fmt.Println("Instalando dependencia: s-lang")
exec.Command("latte", "install", "s-lang")
	fmt.Println("Instalando dependencia: unac")
exec.Command("latte", "install", "unac")
}
