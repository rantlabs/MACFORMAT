# MACFORMAT

Single, small batch, or bulk mac address formating executable command line application. MACFORMAT outputs
all three MAC address formats and appends them to the input line

```
macformat -h
Usage of macformat:
  -file string
    	File containing MAC addresses
  -mac string
    	Single MAC address to format
  -output string
    	Output file
```

## Several input methods

= Input via STDIN (Standard Input) - Unix pipe or STDIN redirect (<) a file with MAC addresses or any input with MAC addresses.
= Input via command line options shown above -mac or -file.

## Examples

```
echo "841b.5e91.7e2d" | macformat 
Output includes original line plus all three differenct MAC address formats appended to the end of the line.
841b.5e91.7e2d DASH: 84-1B-5E-91-7E-2D DOT: 841B.5E91.7E2D COLON: 84:1B:5E:91:7E:2D

Example 2: Inputing data via unix pipes and redirects.
File (file.txt) with the following mac address:
1. 9cc9.ebd2.3b9d
2. 9418.65e9.f240
3. 5407.7d1c.a2f9
4. 6cb0.ce22.b7c3

macformat < file.txt 
1. 9cc9.ebd2.3b9d DASH: 9C-C9-EB-D2-3B-9D DOT: 9CC9.EBD2.3B9D COLON: 9C:C9:EB:D2:3B:9D
2. 9418.65e9.f240 DASH: 94-18-65-E9-F2-40 DOT: 9418.65E9.F240 COLON: 94:18:65:E9:F2:40
3. 5407.7d1c.a2f9 DASH: 54-07-7D-1C-A2-F9 DOT: 5407.7D1C.A2F9 COLON: 54:07:7D:1C:A2:F9
4. 6cb0.ce22.b7c3 DASH: 6C-B0-CE-22-B7-C3 DOT: 6CB0.CE22.B7C3 COLON: 6C:B0:CE:22:B7:C3

cat file.txt| macformat
1. 9cc9.ebd2.3b9d DASH: 9C-C9-EB-D2-3B-9D DOT: 9CC9.EBD2.3B9D COLON: 9C:C9:EB:D2:3B:9D
2. 9418.65e9.f240 DASH: 94-18-65-E9-F2-40 DOT: 9418.65E9.F240 COLON: 94:18:65:E9:F2:40
3. 5407.7d1c.a2f9 DASH: 54-07-7D-1C-A2-F9 DOT: 5407.7D1C.A2F9 COLON: 54:07:7D:1C:A2:F9
4. 6cb0.ce22.b7c3 DASH: 6C-B0-CE-22-B7-C3 DOT: 6CB0.CE22.B7C3 COLON: 6C:B0:CE:22:B7:C3

Example 3: Using the command line application flags
macformat -file file.txt
1. 9cc9.ebd2.3b9d DASH: 9C-C9-EB-D2-3B-9D DOT: 9CC9.EBD2.3B9D COLON: 9C:C9:EB:D2:3B:9D
2. 9418.65e9.f240 DASH: 94-18-65-E9-F2-40 DOT: 9418.65E9.F240 COLON: 94:18:65:E9:F2:40
3. 5407.7d1c.a2f9 DASH: 54-07-7D-1C-A2-F9 DOT: 5407.7D1C.A2F9 COLON: 54:07:7D:1C:A2:F9
4. 6cb0.ce22.b7c3 DASH: 6C-B0-CE-22-B7-C3 DOT: 6CB0.CE22.B7C3 COLON: 6C:B0:CE:22:B7:C3

macformat -mac "841b.5e91.7e2d"
DASH: 84-1B-5E-91-7E-2D DOT: 841B.5E91.7E2D COLON: 84:1B:5E:91:7E:2D

macformat -mac 84-1b-5e-91-7e-2d
DASH: 84-1B-5E-91-7E-2D DOT: 841B.5E91.7E2D COLON: 84:1B:5E:91:7E:2D
```

## Output

    - Output is STDOUT or the terminal screen. If the -output flag is enabled, provide a file name.

## Speed

macformat reformat thousands of mac address per second

## Download

Download the appropriate executable for your device from this repo and just run the executable with no arguments. You should see the entire help message show above.

- macformatv4_linux_32 Linux 32 bit
- macformatv4_linux_64 Linux 64 bit
- macformatv4_mac Apple
- macformatv4_mac_arm64 Apple Silicon
- macformatv4_rpi_arm64 ARM64 (used in Raspberry Pi 3 Model B+ and 4)
- macformatv4_rpi_armv6 Raspberry Pi (ARM architecture)
- macformatv4_rpi_armv7 For ARMv7 (used in Raspberry Pi 2 and 3)
- macformatv4_windows_32.exe Windows 32 bit
- macformatv4_windows_64.exe Windows 64 bit
